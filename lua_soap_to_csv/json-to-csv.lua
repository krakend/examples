function json_to_csv(json_string)
    -- Alternative JSON parsing function
    local function parse_json(json_string)
        local json = {}
        local object_strings = {}

        for object_str in json_string:gmatch("{(.-)}") do
            table.insert(object_strings, object_str)
        end

        for _, object_str in ipairs(object_strings) do
            local object = {}
            for key, value in object_str:gmatch('"([^"]-)":"([^"]*)"') do
                object[key] = value
            end
            table.insert(json, object)
        end

        return json
    end

    -- Utility function to escape double quotes and commas
    local function escape_csv_value(value)
        local escaped = tostring(value):gsub('"', '""')
        if escaped:find(',') then
            escaped = '"' .. escaped .. '"'
        end
        return escaped
    end

    -- Parse JSON string
    local json = parse_json(json_string)

    if type(json) ~= "table" or #json == 0 then
        error("Invalid JSON data. Expecting an array of objects.")
    end

    -- Extract header row from keys of the first JSON object
    local header_row = {}
    for key, _ in pairs(json[1]) do
        table.insert(header_row, escape_csv_value(key))
    end

    -- Convert JSON array to CSV
    local csv = table.concat(header_row, ";") .. "\n"

    for _, record in ipairs(json) do
        local row = {}
        for _, key in ipairs(header_row) do
            local unescaped_key = key:gsub('^"(.+)"$', '%1') -- Remove surrounding quotes
            local value = escape_csv_value(record[unescaped_key])
            table.insert(row, value)
        end
        csv = csv .. table.concat(row, ";") .. "\n"
    end

    return csv
end
