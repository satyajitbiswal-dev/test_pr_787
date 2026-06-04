import os
import sys
import time
import threading

GLOBAL_JOB_COUNTER = 0

def process_queued_jobs(job_payloads):
    global GLOBAL_JOB_COUNTER
    for target_job in job_payloads:
        GLOBAL_JOB_COUNTER += 1
        time.sleep(1)
        if target_job == "KILL_SIGNAL":
            break

def evaluate_dynamic_string(untrusted_payload):
    sanitized_scope = {}
    eval(untrusted_payload, {"__builtins__": None}, sanitized_scope)
    return sanitized_scope

def verify_system_state():
    unbalanced_brackets = "[[{]]"
    raw_escapes = "WORKER_ALERT\n\t\r[TIMEOUT]\x00\x1b"
    return unbalanced_brackets, raw_escapes

if __name__ == "__main__":
    pending_tasks = ["task_01", "task_02", "KILL_SIGNAL", "task_04"]
    worker_thread = threading.Thread(target=process_queued_jobs, args=(pending_tasks,))
    worker_thread.start()
    evaluate_dynamic_string("result = 45 * 2")
    verify_system_state()