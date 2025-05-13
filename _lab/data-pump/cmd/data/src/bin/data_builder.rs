#![allow(non_snake_case)]

use chrono::{DateTime, Local};
use data::config::DATA_COUNT;
use rand::Rng;
use serde::Serialize;
use std::fs::{self, File};
use std::io::Write;
use uuid::Uuid;

#[derive(Serialize)]
struct Owner {
    id: String,
    owner: String,
    owner__first_name: String,
}

#[derive(Serialize)]
struct Property {
    id: String,
    property_id: String,
    parent: Option<String>,
    children: Vec<String>,
    block_no__code: String,
    resurvey_no: String,
    survey_type: String,
    sub_division_no: String,
    area: f64,
    #[serde(rename = "type")]
    type_field: String,
    classification: String,
    basic_tax_rate: f64,
    owned_date: String,
    forfeited_date: Option<String>,
    is_freezed: bool,
}

#[derive(Serialize)]
struct Data {
    id: String,
    tp_no: String,
    village: u32,
    village__name: String,
    village__parent__name: String,
    village__parent__district__name: String,
    ownership: Vec<Owner>,
    created_at: DateTime<Local>,
    properties: Vec<Property>,
}

fn random_owner() -> Owner {
    Owner {
        id: Uuid::new_v4().to_string(),
        owner: Uuid::new_v4().to_string(),
        owner__first_name: "John".to_string(),
    }
}

fn random_property() -> Property {
    let now = Local::now();
    let area = rand::rng().random_range(100.0..1000.0);
    Property {
        id: Uuid::new_v4().to_string(),
        property_id: format!(
            "628535|N|965|{}|23|{:.4}|1|{}|{}",
            rand::rng().random_range(100000..999999),
            area,
            now.format("%Y-%m-%d"),
            now.format("%H:%M:%S")
        ),
        parent: None,
        children: vec![],
        block_no__code: "965".to_string(),
        resurvey_no: rand::rng().random_range(100000..999999).to_string(),
        survey_type: "new".to_string(),
        sub_division_no: "23".to_string(),
        area,
        type_field: "test".to_string(),
        classification: "land".to_string(),
        basic_tax_rate: 25.0,
        owned_date: now.format("%Y-%m-%d").to_string(),
        forfeited_date: None,
        is_freezed: false,
    }
}

fn generate_data(tp: i32) -> Data {
    let now = Local::now();
    Data {
        id: Uuid::new_v4().to_string(),
        tp_no: tp.to_string(),
        village: 182,
        village__name: "Thirupuram".to_string(),
        village__parent__name: "Neyyattinkara".to_string(),
        village__parent__district__name: "THIRUVANANTHAPURAM".to_string(),
        ownership: vec![random_owner()],
        created_at: now,
        properties: vec![random_property()],
    }
}

fn main() {
    let data: Vec<Data> = (0..DATA_COUNT).map(|i| generate_data(i + 1)).collect();
    let json = serde_json::to_string_pretty(&data).unwrap();

    let _ = fs::create_dir_all("./dist").expect("Could not create folder");
    let mut file = File::create("./dist/dummy_data.json").expect("Could not create file");
    file.write_all(json.as_bytes())
        .expect("Could not write data");

    println!("dummy_data.json created with {} entries", DATA_COUNT);
}
