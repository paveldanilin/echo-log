require "module"

function filter (text)
    print(">>[" .. text .. "]")
    return true
end

function test()
    return double(400)
end

function process (text) 
    text = text:gsub("%s+", "")
    if text == "we" then
        return 1
    end 
    if text == "wer" then
        return 2
    end
    return 0
end

function print_table(table) 
    for k, v in pairs(table) do
        if type(v) == "table" then
            print(k)
            print_table(v)
        else
            print(k, v)
        end
    end
end

function check_access(person, flag, n)
    -- print("<<")
    -- print("=>" .. tostring(flag)) 
    -- print("=>" .. tostring(n))
    print_table(person)
    -- print(">>")
    if person["age"] >= 30 then
        print("OLD")
        return test_one()
    elseif person["age"] < 30 then
        print("YOUNGBLOOD")
        return test_two()
    end
end