const std = @import("std");
const one = @import("days/one.zig");
const two = @import("days/two.zig");
const three = @import("days/three.zig");
const four = @import("days/four.zig");
// const five = @import("days/five.zig");
// const six = @import("days/six.zig");
// const seven = @import("days/seven.zig");
// const eight = @import("days/eight.zig");

const TEN_MB = 1024 * 1000 * 10;

pub fn run(allocator: std.mem.Allocator, day: u8, test_data: bool) !void {
    // turn day num into number_str
    var b: [64]u8 = undefined;
    var day_str: []const u8 = undefined;
    switch (day) {
        1 => day_str = "one",
        2 => day_str = "two",
        3 => day_str = "three",
        4 => day_str = "four",
        5 => day_str = "five",
        6 => day_str = "six",
        7 => day_str = "seven",
        8 => day_str = "eight",
        else => unreachable,
    }

    // load the data
    const fp = try std.fmt.bufPrint(&b, "../data/{s}.txt", .{day_str});
    std.debug.print("Loading: {s}\n", .{fp});
    var iter = try get_intput(allocator, fp);
    defer allocator.free(iter.buffer);
    var input = iter.next().?;
    if (!test_data) {
        input = iter.next().?;
    }
    if (input[0] == '\n') {
        input = input[1..]; // skip opening newline
    }
    if (input[input.len - 1] == '\n') {
        input = input[0 .. input.len - 1];
    }

    // run the code
    switch (day) {
        1 => try one.run(input),
        2 => try two.run(allocator, input),
        3 => try three.run(input),
        4 => try four.run(allocator,input),
        else => try one.run(input),
    }
}

fn get_intput(allocator: std.mem.Allocator, fp: []const u8) !std.mem.SplitIterator(u8, .sequence) {
    const f = try std.fs.cwd().readFileAlloc(allocator, fp, TEN_MB);
    return std.mem.splitSequence(u8, f, "--**--**-It's the AOC-**--**--");
}
// pub fn bufferedPrint() !void {
//     // Stdout is for the actual output of your application, for example if you
//     // are implementing gzip, then only the compressed bytes should be sent to
//     // stdout, not any debugging messages.
//     var stdout_buffer: [1024]u8 = undefined;
//     var stdout_writer = std.fs.File.stdout().writer(&stdout_buffer);
//     const stdout = &stdout_writer.interface;
//
//     try stdout.print("Run `zig build test` to run the tests.\n", .{});
//
//     try stdout.flush(); // Don't forget to flush!
// }
//
// pub fn add(a: i32, b: i32) i32 {
//     return a + b;
// }
//
// test "basic add functionality" {
//     try std.testing.expect(add(3, 7) == 10);
// }
