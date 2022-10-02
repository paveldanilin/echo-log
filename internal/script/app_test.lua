-- Load module
local acc = require "counter_test"

function main()
    a = acc.new(999)
    a:inc()
    return a:get_value()
end