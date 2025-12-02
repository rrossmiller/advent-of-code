const std = @import("std");
const TEN_MB = 1024 * 1000 * 10;

pub fn run(input: []const u8) !void {
    std.debug.print("Running day TWO\n", .{});
    try p1(input);
    // try p2(input);
}

fn p1(input: []const u8) !void {
    std.debug.print("p1: {s}\n", .{input});
}
// fn p2(input: []const u8) !void {
//     //
// }

fn get_intput(allocator: std.mem.Allocator) ![]const u8 {
    if (true) {
        const x =
            \\two
        ;
        return try allocator.dupe(u8, x);
    }
    const f = try std.fs.cwd().readFileAlloc(allocator, "../data/two.txt", TEN_MB);
    return f;
}
