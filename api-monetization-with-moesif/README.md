# KrakenD with Moesif Integration Example

## Overview

This repository showcases the integration of KrakenD with Moesif, providing advanced API analytics and monetization capabilities. Moesif is a robust analytics and billing platform, which, when combined with KrakenD, facilitates the monitoring, governance, and monetization of API usage.

## Features

- **Moesif Integration:** Seamless sending of API events to Moesif for detailed analytics.
- **Real-time Synchronization:** KrakenD is synchronized with Moesif for real-time governance rule application.
- **Enhanced Monitoring:** Track API usage by customers, manage subscriptions, and enforce quotas.
- **Effortless API Monetization:** Connect to payment gateways and set up usage-based billing.

## Prerequisites

- Docker
- KrakenD Enterprise Edition (EE) version 2.5.0 or later

## Configuration

The `krakend.json` in this repository is configured for integration with Moesif. Key configurations include:

- **Application ID:** Unique identifier for Moesif application.
- **User and Company Identification:** Configure headers or JWT claims to identify users and companies.
- **Event Queue and Batch Size:** Manage memory and CPU consumption efficiently.
- **Debug and Logging Options:** Enhanced logging for development and production environments.

## Running KrakenD with Moesif Integration

Run KrakenD with Moesif integration using Docker:

```bash
docker run -p "8080:8080" -v "$PWD:/etc/krakend" krakend/krakend-ee:2.5
```

## Example Usage

Test the integration with a sample request:

```bash
curl -H 'X-Tenant: Customer5' -H 'X-Company: YourCompany' 'http://localhost:8080/test?key=microsoft'

```

## Monitoring and Analytics

- KrakenD logs events related to Moesif, including middleware injection and event queuing.
- Use Moesif's dashboard for in-depth analysis of API usage and performance.

## Key Benefits

- **Analyze Customer API Usage:** Real-time insights into API usage patterns.
- **Ensure Customer Success:** Automated notifications and guidance for developers.
- **Charge for APIs:** Set up API-based billing with ease.
- **Enforce Subscription Quotas:** Automatically manage freemium and trial accounts.

## Configuration Details

- **`application_id`**: Moesif Collector Application ID.
- **`batch_size`**: Number of events sent in each batch.
- **`event_queue_size`**: In-memory event queue size.
- **`identify_company`**: Methods to identify companies (header, JWT claim, query string).

For detailed configuration options, refer to the [Moesif Schema](https://www.krakend.io/schema/v2.5/telemetry/moesif.json).
