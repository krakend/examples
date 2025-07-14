# Workflows example

This example simulates a backend for frontend example for a streaming 
platforme like could be Nerdflix, or Rainforest Prime Video.

A customer can have several profiles. 

## Endpoints

### `/home_dashboard`

Is modeled as the initial screen that a customer would watch.

The request can specify the `customer` and optionally the `profile`. 
If the profile is not provided, it would request the default profile for 
a given customer in a sequential workflow.

We have three backends to request in parallel:

- header
- customized content
- footer

**Customized content** is a workflow, because if `profile` is not provided we want
to have sequential call that first fetches the default profile. Once we are
sure we have `profile` available, we can call another workflow that will 
execute the fetch of customized content from different backends.
