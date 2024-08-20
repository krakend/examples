function merge_reviewers(resp)
    local response_body = resp:data()
    local reviews = response_body:get("reviews")
    local user_ids = {}

    -- Get users' IDs from the reviews data.
    for i = 1, reviews:len() do
        local review = reviews:get(i - 1)
        table.insert(user_ids, review:get("userId"))
    end
    local user_ids_str = table.concat(user_ids, ",")

    -- Retrieve users' details from the user microservice
    local user_response = http_response.new("http://my_service/users.json?ids=" .. user_ids_str)

    if user_response:statusCode() == 200 then
        local user_data_str = user_response:body()
        local user_data = json_parse(user_data_str)

        -- Add user details to the reviews
        for i = 1, reviews:len() do
            local review = reviews:get(i - 1)
            local user_id = review:get("userId")
            for _, user in ipairs(user_data) do
                if user.id == user_id then
                    review:set("userName", user.name)
                    review:set("userEmail", user.email)
                    review:set("userProfilePicture", user.profilePicture)
                end
            end
        end
    else
        print("Failed to fetch user data")
    end
end
