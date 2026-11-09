# From SOAP to REST with multiple formats 
This KrakenD configuration exposes three endpoints, each presenting continent data in different formats: JSON, XML, and CSV. Origin data is collected from a legacy SOAP service.

### Endpoints:
1. `/continents.json`
   This endpoint retrieves continent data from the CountryInfoService SOAP web service at http://webservices.oorsprong.org/websamples.countryinfo/CountryInfoService.wso. The data is fetched with a POST request and converted from XML to a JSON object using a JMESPath expression. The output is a JSON collection.

2. `/continents.xml`
   This endpoint fetches the JSON data from the /continents.json endpoint and converts it to an XML format. The output is an XML document with a root element called continents containing one continent element for each continent.

3. `/continents.csv`
   This endpoint retrieves the JSON data from the /continents.json endpoint and uses a Lua script (./json-to-csv.lua) to convert the JSON object into a CSV format. The output is a CSV file with a Content-Type header set to text/csv.

### Usage
This example needs KrakenD Enterprise, because the SOAP integration is an Enterprise feature. Replace the placeholder `LICENSE` file in this folder with yours, and start KrakenD from this folder:

```bash
docker run --rm -p 8080:8080 -v "$PWD:/etc/krakend" krakend/krakend-ee:3.0 run -dc krakend.json
```

The backend is the public [CountryInfoService](http://webservices.oorsprong.org/websamples.countryinfo/CountryInfoService.wso), so you need internet access.

You can query the endpoints using any HTTP client, such as curl or a web browser, to obtain the data in the desired format:

- JSON: http://localhost:8080/continents.json
- XML: http://localhost:8080/continents.xml
- CSV: http://localhost:8080/continents.csv
