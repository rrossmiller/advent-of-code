const std = @import("std");

pub fn run(input: []const u8) !void {
    std.debug.print("Running day ONE\n", .{});
    // std.debug.print("*{d}\n", .{@mod(-1012, 100)});
    // std.debug.print("*{d}\n", .{-1012 / 100});
    // std.debug.print("\n", .{});
    // std.debug.print("*{d}\n", .{@mod(1012, 100)});
    // std.debug.print("*{d}\n", .{1012 / 100});
    // std.debug.print("*{d}\n", .{112 / 100});
    // if (true)
    //     return;
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
    std.debug.print("p1: {d}\n", .{zero});
}

fn p2(input: []const u8) !void {
    var dial: i32 = 50;
    var zero: i32 = 0;
    var lines = std.mem.splitScalar(u8, input, '\n');
    while (lines.next()) |val| {
        if (val.len == 0) {
            break;
        }

        const was_zero = dial == 0;
        var had_update = false;

        const dir = val[0];
        const i = try std.fmt.parseInt(i32, val[1..], 10);
        std.debug.print("{d} ({c}) {d}", .{ dial, dir, i });
        switch (dir) {
            'L' => {
                dial -= i;
                std.debug.print(" -> {d}", .{dial});
                if (dial == 0) zero += 1;
                while (dial < 0) {
                    dial = 100 + dial;
                    zero += 1;
                    had_update = true;
                }

                if (was_zero and had_update) zero -= 1;
                std.debug.print(" ({d}) : {d} ({})\n", .{ dial, zero, was_zero });
            },
            'R' => {
                dial += i;
                std.debug.print(" -> {d}", .{dial});
                if (dial == 0) zero += 1;
                while (dial > 99) {
                    dial = dial - 100;
                    zero += 1;
                    had_update = true;
                }

                if (was_zero and had_update) zero -= 1;
                std.debug.print(" ({d}) : {d} ({})\n", .{ dial, zero, was_zero });
            },
            else => unreachable,
        }
    }
    std.debug.print("p2: {d}\n", .{zero});
    std.debug.print("6273 too high\n", .{});
    std.debug.print("6070 not right\n", .{});
}

/// doesn't work
fn p2_mod(input: []const u8) !void {
    var dial: i32 = 50;
    var zero: i32 = 0;
    var lines = std.mem.splitScalar(u8, input, '\n');
    while (lines.next()) |val| {
        if (val.len == 0) {
            break;
        }

        // don't count going past zero if we start at 0
        var was_zero = false;
        if (dial == 0) {
            was_zero = true;
        }

        const dir = val[0];
        const i = try std.fmt.parseInt(i32, val[1..], 10);
        std.debug.print("{d} ({c}) {d}", .{ dial, dir, i });
        switch (dir) {
            'L' => {
                dial -= i;
            },
            'R' => {
                dial += i;
            },
            else => unreachable,
        }
        std.debug.print("-> {d}", .{dial});

        if (dial == 100 or dial == 0) {
            dial = 0;
            zero += if (!was_zero) 1 else 0;
        } else if (dial < -99) {
            var times = @divFloor(dial, 100);
            if (times < 0) times *= -1;
            if (was_zero) {
                times -= 1;
            }
            zero += times;
            std.debug.print(" - {d} times {} -", .{ times, was_zero });

            dial = @mod(dial, 100);
        } else if (dial < 0) {
            dial = @mod(dial, 100);
            // zero += 1;
            zero += if (!was_zero) 1 else 0;
        } else if (dial > 100) {
            var times = @divFloor(dial, 100);
            if (was_zero) {
                times -= 1;
            }
            zero += times;
            std.debug.print(" - {d} times {} -", .{ times, was_zero });

            dial = @mod(dial, 100);
        }
        std.debug.print(" : {d} {d}\n", .{ dial, zero });
    }
    std.debug.print("p2: {d}\n", .{zero});
    std.debug.print("6273 too high\n", .{});
    std.debug.print("6070 not right\n", .{});
}
