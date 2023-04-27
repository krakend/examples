local function get_items_from_page(users_list, page, items_per_page)
    local start_index = (page - 1) * items_per_page + 1
    local end_index = start_index + items_per_page - 1

    local result = {}
    for i = start_index, end_index do
        if users_list[i] ~= nil then
            table.insert(result, users_list[i])
        end
    end
    return result
end

local function parse_json_list(json_string, list_name)
    local list_start_pattern = '"' .. list_name .. '":%['
    local start_index = string.find(json_string, list_start_pattern)

    if not start_index then
        return {}
    end

    local list_start = start_index + #list_start_pattern - 1
    local list_end = string.find(json_string, "%]", list_start)
    local list_content = string.sub(json_string, list_start, list_end - 1)

    local items = {}
    for item in string.gmatch(list_content, "%b{}") do
        table.insert(items, item)
    end

    return items
end

function merge_users_and_paginate(req, resp)
    local page = tonumber(req:params('Page')) or 1
    local items_per_page = 10

    local response_body = resp:body()

    local users_list = parse_json_list(response_body, "users_list")

    local paged_users = get_items_from_page(users_list, page, items_per_page)

    return "[" .. table.concat(paged_users, ",") .. "]"
end
