# Workflows example

This example simulates a backend for frontend example for a streaming 
platforme like could be Nerdflix, or Rainforest Prime Video.

A customer can have several profiles. 

## Endpoints

### `/sequential_workflow`

Is modeled as the initial screen that a customer would watch.

The request can specify the `customer` and optionally the `profile`. 
If the profile is not provided, it would request the default profile for 
a given customer in a sequential workflow.
