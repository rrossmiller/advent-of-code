const std = @import("std");
const TEN_MB = 1024 * 1000 * 10;

pub fn run() !void {
    // std.debug.print("*{d}\n", .{@mod(1000, 100)});
    // std.debug.print("*{d}\n", .{-1000 / 100});
    // if (true)
    //     return;
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();
    defer {
        const c = gpa.deinit();
        switch (c) {
            .ok => {},
            .leak => std.debug.print("leaked\n", .{}),
        }
    }
    const input = try get_intput(allocator);
    defer allocator.free(input);

    try p1(input);
    try p2(input);
}

fn p1(input: []const u8) !void {
    var dial: i32 = 50;
    var zero: i32 = 0;
    var lines = std.mem.splitScalar(u8, input, '\n');
    while (lines.next()) |val| {
        if (val.len == 0) {
            break;
        }

        const dir = val[0];
        const i = try std.fmt.parseInt(i32, val[1..], 10);
        // std.debug.print("{d} ({c}) {d}: {d}\n", .{ dial, dir, i, zero });
        // std.debug.print("{d}/{d}\n\n", .{ dial + i, dial - i });
        switch (dir) {
            'L' => {
                dial -= i;
                if (dial < 0) {
                    std.debug.print("*{d}\n", .{@mod(dial, 100)});
                }
                while (dial < 0) {
                    dial = 100 + dial;
                }
            },
            'R' => {
                dial += i;
                while (dial > 99) {
                    dial = dial - 100;
                }
            },
            else => unreachable,
        }
        if (dial == 0) {
            zero += 1;
        }
    }
    std.debug.print("{d}\n", .{zero});
    std.debug.print("6376 too high\n", .{});
}
fn p2(input: []const u8) !void {
    var dial: i32 = 50;
    var zero: i32 = 0;
    var lines = std.mem.splitScalar(u8, input, '\n');
    while (lines.next()) |val| {
        if (val.len == 0) {
            break;
        }

        const dir = val[0];
        const i = try std.fmt.parseInt(i32, val[1..], 10);
        std.debug.print("{d} ({c}) {d}: {d}\n", .{ dial, dir, i, zero });
        std.debug.print("{d}/{d}\n\n", .{ dial + i, dial - i });
        switch (dir) {
            'L' => {
                dial -= i;
                if (dial < 0) {
                    std.debug.print("*{d}\n", .{@mod(dial, 100)});
                }
                while (dial < 0) {
                    dial = 100 + dial;
                }
            },
            'R' => {
                dial += i;
                while (dial > 99) {
                    dial = dial - 100;
                }
            },
            else => unreachable,
        }
        if (dial == 0) {
            zero += 1;
        }
    }
    std.debug.print("{d}\n", .{zero});
    std.debug.print("6376 too high\n", .{});
}

fn get_intput(allocator: std.mem.Allocator) ![]const u8 {
    if (true) {
        const x =
            \\L68
            \\L30
            \\R48
            \\L5
            \\R60
            \\L55
            \\L1
            \\L99
            \\R14
            \\L82
        ;
        return try allocator.dupe(u8, x);
    }
    const f = try std.fs.cwd().readFileAlloc(allocator, "../data/one.txt", TEN_MB);
    return f;
}
