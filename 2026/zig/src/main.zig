const std = @import("std");
const aoc = @import("aoc");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();
    defer {
        const c = gpa.deinit();
        switch (c) {
            .ok => {},
            .leak => std.debug.print("leaked\n", .{}),
        }
    }

    var args = try std.process.argsWithAllocator(allocator);
    _ = args.next(); // skip prg name

    var test_data = false;
    var day: u8 = 1;
    while (args.next()) |a| {
        if (a.len < 2) {
            continue;
        }

        if (std.mem.eql(u8, a, "-t") or std.mem.eql(u8, a, "--test")) {
            test_data = true;
        } else if (a[1] >= '1' and a[1] <= '9') {
            day = std.fmt.parseInt(u8, a[1..], 10) catch {
                std.debug.print("Invalid input day: \"{s}\"\n", .{a});
                return;
            };
        }
    }

    errdefer std.debug.print("Ran with test data: {}\n\n", .{test_data});
    try aoc.run(allocator, day, test_data);
    std.debug.print("\nRan with test data: {}\n", .{test_data});
}
