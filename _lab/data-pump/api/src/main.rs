use std::env;

use alloy::{
    providers::{
        Identity, ProviderBuilder, RootProvider,
        fillers::{BlobGasFiller, ChainIdFiller, FillProvider, GasFiller, JoinFill, NonceFiller},
    },
    sol,
};
use axum::{
    Json, Router,
    extract::{Path, State},
    http::StatusCode,
    routing::get,
};
use eyre::Result;
use tower_http::trace::TraceLayer;
use tracing_subscriber::{layer::SubscriberExt, util::SubscriberInitExt};

use Registry::RegistryInstance;

type Instance = RegistryInstance<
    FillProvider<
        JoinFill<
            Identity,
            JoinFill<GasFiller, JoinFill<BlobGasFiller, JoinFill<NonceFiller, ChainIdFiller>>>,
        >,
        RootProvider,
    >,
>;

sol!(
    #[sol(rpc)]
    contract Registry {
    mapping(string => mapping(uint => string)) properties;
    mapping(string => uint) propertyVersions;

    function addProperty(string[] memory landIds, string[] memory data) public {
        for (uint i = 0; i < landIds.length; i++) {
            uint newVersion = propertyVersions[landIds[i]] + 1;
            properties[landIds[i]][newVersion] = data[i];
            propertyVersions[landIds[i]] = newVersion;
        }
    }

    function getLatestProperty(
        string memory propertyId
    ) public view returns (string memory) {
        uint latestVersion = propertyVersions[propertyId];
        return properties[propertyId][latestVersion];
    }

    function getVersionProperty(
        string memory propertyId,
        uint version
    ) public view returns (string memory) {
        return properties[propertyId][version];
    }

    function getPropertyVersions(
        string memory propertyId
    ) public view returns (uint) {
        return propertyVersions[propertyId];
    }

    function getAllPropertyVersions(
        string memory propertyId
    ) public view returns (string[] memory) {
        uint versions = propertyVersions[propertyId];
        string[] memory allVersions = new string[](versions);
        for (uint i = 1; i <= versions; i++) {
            allVersions[i - 1] = properties[propertyId][i];
        }
        return allVersions;
    }
}
);

async fn instance_builder() -> Result<Instance> {
    let chain_url = env::var("CHAIN_URL")?.parse()?;

    let provider = ProviderBuilder::new().connect_http(chain_url);
    let contract = Registry::new(env::var("CONTRACT_ADDRESS")?.parse()?, provider);

    Ok(contract)
}

#[tokio::main]
async fn main() {
    let registry = instance_builder().await.unwrap();

    tracing_subscriber::registry()
        .with(
            tracing_subscriber::EnvFilter::try_from_default_env()
                .unwrap_or_else(|_| "data_pump_api=debug,tower_http=debug".into()),
        )
        .with(tracing_subscriber::fmt::layer())
        .init();

    let listener = tokio::net::TcpListener::bind("0.0.0.0:8090").await.unwrap();

    tracing::debug!("Listening on {}", listener.local_addr().unwrap());
    axum::serve(listener, app(registry)).await.unwrap();
}

fn app(instance: Instance) -> Router {
    let v1_routes = Router::new()
        .route("/ping", get(ping))
        .route("/single/{id}", get(fetch_single))
        .layer(TraceLayer::new_for_http())
        .with_state(instance);

    let api_routes = Router::new().nest("/v1", v1_routes);

    Router::new().nest("/api", api_routes)
}

async fn ping() -> &'static str {
    "pong"
}

async fn fetch_single(
    Path(id): Path<String>,
    State(instance): State<Instance>,
) -> Result<Json<String>, (StatusCode, String)> {
    let result = instance
        .getLatestProperty(id.clone())
        .call()
        .await
        .map_err(internal_error)?;

    Ok(Json(result))
}

fn internal_error<E>(err: E) -> (StatusCode, String)
where
    E: std::error::Error,
{
    (StatusCode::INTERNAL_SERVER_ERROR, err.to_string())
}
