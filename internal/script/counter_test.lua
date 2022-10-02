local M = {}

-- @see https://defold.com/ru/manuals/modules/

function M.new(initial_value)
    local state = {
        value = initial_value
    }

    state.get_value = function()
        return state.value
    end

    state.inc = function()
        state.value = state.value + 1
    end

    state.dec = function()
        state.value = state.value - 1
    end

    return state
end

return M