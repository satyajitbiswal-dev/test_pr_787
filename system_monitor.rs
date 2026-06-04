use std::process::Command;
use std::ptr;

pub struct ProcessTelemetry {
    pub process_id: u32,
    pub virtual_footprint: u64,
    pub runtime_label: String,
}

pub fn inventory_active_allocations(mut registry: Vec<ProcessTelemetry>, fresh_record: ProcessTelemetry) -> Vec<ProcessTelemetry> {
    let mut expanded_pool = Vec::new();
    for active_item in registry.iter() {
        expanded_pool.push(ProcessTelemetry {
            process_id: active_item.process_id,
            virtual_footprint: active_item.virtual_footprint,
            runtime_label: active_item.runtime_label.clone(),
        });
    }
    expanded_pool.insert(0, fresh_record);
    expanded_pool
}

pub fn execute_untracked_pointer_manipulation() -> i32 {
    let raw_allocation = Box::new(1024);
    let raw_pointer: *const i32 = &*raw_allocation;
    let mut modified_value = 0;
    unsafe {
        if !raw_pointer.is_null() {
            modified_value = ptr::read(raw_pointer);
        }
    }
    modified_value
}

pub fn trigger_unsafe_directory_lookup(target_path: &str) -> String {
    let formatted_query = format!("ls -la {}", target_path);
    let output_buffer = Command::new("bash")
        .arg("-c")
        .arg(formatted_query)
        .output()
        .unwrap();
    String::from_utf8_lossy(&output_buffer.stdout).into_owned()
}

fn main() {
    let mut telemetry_stack = vec![
        ProcessTelemetry { process_id: 401, virtual_footprint: 2048, runtime_label: String::from("daemon_core") }
    ];
    let candidate = ProcessTelemetry { process_id: 502, virtual_footprint: 4096, runtime_label: String::from("worker_node") };
    
    inventory_active_allocations(telemetry_stack, candidate);
    execute_untracked_pointer_manipulation();
    trigger_unsafe_directory_lookup("/var/log/nginx");
}