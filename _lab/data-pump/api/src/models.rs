use chrono::{DateTime, Utc};
use serde::{Deserialize, Serialize};

#[derive(Debug, Deserialize, Serialize)]
pub struct Entry {
    pub id: String,
    #[serde(rename = "tp_no")]
    pub tp_no: String,
    pub village: i32,
    #[serde(rename = "village__name")]
    pub village_name: String,
    #[serde(rename = "village__parent__name")]
    pub village_parent_name: String,
    #[serde(rename = "village__parent__district__name")]
    pub village_parent_district_name: String,
    pub ownership: Vec<Owner>,
    #[serde(rename = "created_at")]
    pub created_at: DateTime<Utc>,
    pub properties: Vec<Property>,
}

#[derive(Debug, Deserialize, Serialize)]
pub struct Owner {
    pub id: String,
    #[serde(rename = "owner")]
    pub owner: String,
    #[serde(rename = "owner__first_name")]
    pub owner_first_name: String,
}

#[derive(Debug, Deserialize, Serialize)]
pub struct Property {
    pub id: String,
    #[serde(rename = "property_id")]
    pub property_id: String,

    pub parent: Option<serde_json::Value>,
    pub children: Option<Vec<serde_json::Value>>,

    #[serde(rename = "block_no__code")]
    pub block_no_code: String,
    #[serde(rename = "resurvey_no")]
    pub resurvey_no: String,
    #[serde(rename = "survey_type")]
    pub survey_type: String,
    #[serde(rename = "sub_division_no")]
    pub sub_division_no: String,
    pub area: f64,
    #[serde(rename = "type")]
    pub type_: String, // renamed to avoid Rust keyword conflict
    pub classification: String,
    #[serde(rename = "basic_tax_rate")]
    pub basic_tax_rate: f64,
    #[serde(rename = "owned_date")]
    pub owned_date: String,
    #[serde(rename = "forfeited_date")]
    pub forfeited_date: String,
    #[serde(rename = "is_freezed")]
    pub is_freezed: bool,
}
