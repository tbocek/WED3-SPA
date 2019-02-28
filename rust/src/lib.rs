extern crate wasm_bindgen;

use wasm_bindgen::prelude::*;

#[wasm_bindgen]
pub fn mul(x: i32, y: i32) -> i32 {
    x * y
}

