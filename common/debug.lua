local M = {}

-- https://defold.com/ru/manuals/modules/

function M.new()
   local state = {}

   state.dump = function(o) 
      if type(o) == 'table' then
         local s = '{ '
         for k,v in pairs(o) do
            if type(k) ~= 'number' then k = '"'..k..'"' end
            s = s .. '['..k..'] = ' .. dump(v) .. ','
         end
         return s .. '} '
      else
         return tostring(o)
      end
   end

   return state
end

return M