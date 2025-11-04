function flatten_capital_lat_lng(resp)
    local data = resp:data()
    local country = data:get('rest_data'):get('countries'):get(0)
    local capitalLatLng = country:get('capitalInfo'):get('latlng')

    country:set('capitalLat', tostring(capitalLatLng:get(0)))
    country:set('capitalLng', tostring(capitalLatLng:get(1)))
end