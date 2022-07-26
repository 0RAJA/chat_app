-- 清空所有以k为前缀的键
function DelAllPrefixLua()
    local redisKeys = redis.call('keys', KEYS[1] .. '*');
    for i, k in pairs(redisKeys) do
        redis.call('expire', k, 0);
    end
end
