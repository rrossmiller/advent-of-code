const std = @import("std");

pub fn run(allocator: std.mem.Allocator, input: []const u8) !void {
    std.debug.print("Running day THREE\n", .{});
    try p1(allocator, input);
    // try p2(allocator, input);
}

fn p1(_: std.mem.Allocator, input: []const u8) !void {
    var sum: u32 = 0;

    var iter = std.mem.splitScalar(u8, input, '\n');
    while (iter.next()) |line| {
        // std.debug.print("{s}\n", .{line});
        // track max
        var max: u32 = 0;
        for (0..line.len) |l| {
            const ln = try std.fmt.parseInt(u32, line[l .. l + 1], 10);
            // skip to next l if l  < int(max / 10)
            // std.debug.print("{d} {d} = {d} ({})\n", .{ ln, max, max / 10 ,ln});
            if (ln < max / 10) {
                continue;
            }
            for (l + 1..line.len) |r| {
                const rn = try std.fmt.parseInt(u32, line[r .. r + 1], 10);
                const n = ln * 10 + rn;
                // std.debug.print("{d} - ", .{n});
                if (n > max) max = n;
            }
            // std.debug.print("\n", .{});
        }
        // std.debug.print("{d}\n\n", .{max});

        sum += max;
    }
    std.debug.print("p1: {d}\n", .{sum});
}
