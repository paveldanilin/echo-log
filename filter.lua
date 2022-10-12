function lw_filter_event(e)
  -- ntf:Notify("bzz", "lua", e:GetField("email"))
  -- print(e:GetField("email"))
  -- return tonumber(e:GetField("age")) > 20
  if e:GetField("loglevel") == "ERROR" then
    return true
  end
  return false
end
