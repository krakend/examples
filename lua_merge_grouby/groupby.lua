function parse_and_group_json(json_string, group_key)
    local json = {}
    local grouped_items = {}

    json_string = json_string:match("%[(.*)%]") -- Extract the contents of the array

    for object_str in json_string:gmatch("{(.-)}") do
        local object = {}
        for key, value in object_str:gmatch('"([^"]-)":("?[^",}]+[^",}]?"?)') do
            if value:match("^\"") and value:match("\"$") then
                value = value:sub(2, -2) -- Remove quotes if the value is a string
            elseif tonumber(value) then
                value = tonumber(value) -- Convert to number if the value is numeric
            end
            object[key] = value
        end

        local group = object[group_key]
        if not grouped_items[group] then
            grouped_items[group] = object
        else
            for k, v in pairs(object) do
                if k ~= group_key then
                    grouped_items[group][k] = v
                end
            end
        end
    end

    for _, item in pairs(grouped_items) do
        table.insert(json, item)
    end

    return json
end

function table_to_json(val, indent)
    indent = indent or 0
    local str = {}
    local function convert(val)
        if type(val) == "table" then
            if #val > 0 then
                table.insert(str, "[")
                for i, v in ipairs(val) do
                    if i > 1 then table.insert(str, ",") end
                    table.insert(str, "\n" .. string.rep("  ", indent + 1))
                    indent = indent + 1
                    convert(v)
                    indent = indent - 1
                end
                table.insert(str, "\n" .. string.rep("  ", indent) .. "]")
            else
                table.insert(str, "{")
                local first = true
                for k, v in pairs(val) do
                    if not first then table.insert(str, ",") end
                    table.insert(str, "\n" .. string.rep("  ", indent + 1) .. "\"" .. k .. "\": ")
                    indent = indent + 1
                    convert(v)
                    indent = indent - 1
                    first = false
                end
                table.insert(str, "\n" .. string.rep("  ", indent) .. "}")
            end
        elseif type(val) == "string" then
            table.insert(str, "\"" .. val .. "\"")
        else
            table.insert(str, tostring(val))
        end
    end
    convert(val)
    return table.concat(str)
end

function groupby(resp)
    local response_body = resp:body()
    local grouped_components_list = parse_and_group_json(response_body, 'Component No')
    return resp:body(table_to_json(grouped_components_list))
end
