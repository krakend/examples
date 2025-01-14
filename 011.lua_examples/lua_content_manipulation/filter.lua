function filter(req, resp)
    local response_data = resp:data()

    -- remove all existing keys
    response_data:del('bomCount');
    response_data:del('bomList');
    response_data:del('componentCount');
    response_data:del('componentList');
    response_data:del('itemCount');
    response_data:del('itemList');
    response_data:del('referenceDesignators');

     --get all keys (next rele
    --local keys = response_data:keys()

     --iterate over all keys to remove them
    --for _, key in ipairs(keys) do
    --    response_data:del(key)
    --end
end
