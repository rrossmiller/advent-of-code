const std = @import("std");
const TEN_MB = 1024 * 1000 * 10;

pub fn run(allocator: std.mem.Allocator, input: []const u8) !void {
    std.debug.print("Running day TWO\n", .{});
    try p1(allocator, input);
    try p2(allocator, input);
}

fn p1(allocator: std.mem.Allocator, input: []const u8) !void {
    var iter = std.mem.splitScalar(u8, input, ',');

    const digit_string_buf: []u8 = try allocator.alloc(u8, 256);
    defer allocator.free(digit_string_buf);
    var ttl_invalid_ids: usize = 0;
    while (iter.next()) |r| {
        // std.debug.print("{s}\n", .{r});
        var rng_iter = std.mem.splitScalar(u8, r, '-');
        const rng_start = try std.fmt.parseInt(usize, rng_iter.next().?, 10);
        const rng_end = try std.fmt.parseInt(usize, rng_iter.next().?, 10);

        var invalid_ids: usize = 0;
        for (rng_start..rng_end + 1) |i| {
            const digit_string = try std.fmt.bufPrint(digit_string_buf, "{d}", .{i});
            // only ids with even number of digits can be invalid
            if (digit_string.len % 2 == 0) {
                // std.debug.print("{d},", .{i});
                // std.debug.print("({s})", .{digit_string[0 .. digit_string.len / 2]});
                // check first half and second half
                if (std.mem.eql(u8, digit_string[0 .. digit_string.len / 2], digit_string[digit_string.len / 2 ..])) {
                    invalid_ids += i;
                }
            }
        }
        // std.debug.print("\n\n", .{});
        ttl_invalid_ids += invalid_ids;
    }
    std.debug.print("p1: {d}\n", .{ttl_invalid_ids});
}
fn p2(allocator: std.mem.Allocator, input: []const u8) !void {
    if (true)
        return;

    var iter = std.mem.splitScalar(u8, input, ',');

    const digit_string_buf: []u8 = try allocator.alloc(u8, 256);
    defer allocator.free(digit_string_buf);
    var ttl_invalid_ids: usize = 0;
    while (iter.next()) |r| {
        // std.debug.print("{s}\n", .{r});
        var rng_iter = std.mem.splitScalar(u8, r, '-');
        const rng_start = try std.fmt.parseInt(usize, rng_iter.next().?, 10);
        const rng_end = try std.fmt.parseInt(usize, rng_iter.next().?, 10);

        var invalid_ids: usize = 0;
        // for every number in the range
        for (rng_start..rng_end + 1) |i| {
            const digit_string = try std.fmt.bufPrint(digit_string_buf, "{d}", .{i});
            std.debug.print("{s}\n", .{digit_string});
            // take windows of len 1 - digit_string.length/2+1
            for (1..digit_string.len / 2 + 1) |win_len| {
                // can't repeat if the windows don't have the same number of chars
                if (@mod(digit_string.len, win_len) != 0) {
                    continue;
                }

                const w_0 = digit_string[0..win_len];
                const n_divs = @divFloor(digit_string.len, win_len);
                std.debug.print("({s} len={d}, win_len={d}), /={d}  | ", .{ w_0, digit_string.len, win_len, @divFloor(digit_string.len, win_len) });
                var idx: usize = 1;
                var match = true;
                // compare if all windows are the same
                while (idx < n_divs) : (idx += 1) {
                    const w = digit_string[idx * win_len .. (idx + 1) * win_len];

                    std.debug.print(" ?{s} {s} {} ({d})? ", .{ w_0, w, std.mem.eql(u8, w_0, w), idx });
                    // if they don't match break to next number
                    if (!std.mem.eql(u8, w_0, w)) {
                        match = false;
                        break;
                    }
                }

                // if all windows are the same, increas invalid_ids
                if (match) {
                    std.debug.print(" < = {d} >", .{i});
                    invalid_ids += i;
                    // skip next window scanning - advance to the next number
                    break;
                }
            }
            std.debug.print("\n", .{});
        }
        std.debug.print("\n\n", .{});
        ttl_invalid_ids += invalid_ids;
    }
    std.debug.print("p2: {d}\n", .{ttl_invalid_ids});
    std.debug.print("L:  4174379265\n", .{});
    // std.debug.print("t:  {d}\n", .{4174379265});
}
