const std = @import("std");

pub fn run(allocator: std.mem.Allocator, input: []const u8) !void {
    std.debug.print("Running day FOUR\n", .{});
    try p1(allocator, input);
    try p2(allocator, input);
}

fn p1(allocator: std.mem.Allocator, input: []const u8) !void {
    var rolls: i32 = 0;

    var grid = try std.ArrayList(std.ArrayList(u8)).initCapacity(allocator, 100);

    var rows = std.mem.splitScalar(u8, input, '\n');
    while (rows.next()) |row| {
        // insert row
        var g = try std.ArrayList(u8).initCapacity(allocator, 512);
        try g.appendSlice(allocator, row);
        try grid.append(allocator, g);
        // std.debug.print("{s}\n", .{g.items});
    }
    const nrows = grid.items.len;
    const ncols = grid.items[0].items.len;

    // find rolls with < 4 neighbors
    for (0..nrows) |row| {
        for (0..ncols) |col| {
            if (grid.items[row].items[col] == '@') {
                // look around for other @
                var n: i32 = 0;

                // up left
                var y = if (row > 0) row - 1 else nrows;
                var x = if (col > 0) col - 1 else ncols;
                if (y >= 0 and y < nrows and x >= 0 and x < ncols and grid.items[y].items[x] == '@') {
                    n += 1;
                }

                //up
                x = col;
                if (y >= 0 and y < nrows and x >= 0 and x < ncols and grid.items[y].items[x] == '@') {
                    n += 1;
                }
                // up right
                x = col + 1;
                if (y >= 0 and y < nrows and x >= 0 and x < ncols and grid.items[y].items[x] == '@') {
                    n += 1;
                }
                // left
                y = row;
                x = if (col > 0) col - 1 else ncols;
                if (y >= 0 and y < nrows and x >= 0 and x < ncols and grid.items[y].items[x] == '@') {
                    n += 1;
                }

                // right
                x = col + 1;
                if (y >= 0 and y < nrows and x >= 0 and x < ncols and grid.items[y].items[x] == '@') {
                    n += 1;
                }

                // down left
                y = row + 1;
                x = if (col > 0) col - 1 else ncols;
                if (y >= 0 and y < nrows and x >= 0 and x < ncols and grid.items[y].items[x] == '@') {
                    n += 1;
                }

                // down
                x = col;
                if (y >= 0 and y < nrows and x >= 0 and x < ncols and grid.items[y].items[x] == '@') {
                    n += 1;
                }

                // down right
                x = col + 1;
                if (y >= 0 and y < nrows and x >= 0 and x < ncols and grid.items[y].items[x] == '@') {
                    n += 1;
                }

                if (n < 4) {
                    // std.debug.print("{d}, {d} = {d}\n", .{ row, col, n });
                    // grid.items[row].items[col] = 'X';
                    rolls += 1;
                    // break;
                }
            }
        }
    }

    // std.debug.print("\n", .{});
    // for (0..nrows) |row| {
    //     const a = grid.items[row].items;
    //     std.debug.print("{s}\n", .{a});
    // }

    // cleanup
    for (0..grid.items.len) |i| {
        grid.items[i].deinit(allocator);
    }
    grid.deinit(allocator);

    std.debug.print("p1: {d}\n", .{rolls});
}
fn p2(allocator: std.mem.Allocator, input: []const u8) !void {
    var rolls_removed: i32 = 0;
    var rolls: i32 = -1;

    var grid = try std.ArrayList(std.ArrayList(u8)).initCapacity(allocator, 100);

    var rows = std.mem.splitScalar(u8, input, '\n');
    while (rows.next()) |row| {
        // insert row
        var g = try std.ArrayList(u8).initCapacity(allocator, 512);
        try g.appendSlice(allocator, row);
        try grid.append(allocator, g);
        // std.debug.print("{s}\n", .{g.items});
    }
    const nrows = grid.items.len;
    const ncols = grid.items[0].items.len;

    while (rolls != 0) {
        rolls = 0;
        // find rolls with < 4 neighbors
        for (0..nrows) |row| {
            for (0..ncols) |col| {
                if (grid.items[row].items[col] == '@') {
                    // look around for other @
                    var n: i32 = 0;

                    // up left
                    var y = if (row > 0) row - 1 else nrows;
                    var x = if (col > 0) col - 1 else ncols;
                    if (y >= 0 and y < nrows and x >= 0 and x < ncols and grid.items[y].items[x] == '@') {
                        n += 1;
                    }

                    //up
                    x = col;
                    if (y >= 0 and y < nrows and x >= 0 and x < ncols and grid.items[y].items[x] == '@') {
                        n += 1;
                    }
                    // up right
                    x = col + 1;
                    if (y >= 0 and y < nrows and x >= 0 and x < ncols and grid.items[y].items[x] == '@') {
                        n += 1;
                    }
                    // left
                    y = row;
                    x = if (col > 0) col - 1 else ncols;
                    if (y >= 0 and y < nrows and x >= 0 and x < ncols and grid.items[y].items[x] == '@') {
                        n += 1;
                    }

                    // right
                    x = col + 1;
                    if (y >= 0 and y < nrows and x >= 0 and x < ncols and grid.items[y].items[x] == '@') {
                        n += 1;
                    }

                    // down left
                    y = row + 1;
                    x = if (col > 0) col - 1 else ncols;
                    if (y >= 0 and y < nrows and x >= 0 and x < ncols and grid.items[y].items[x] == '@') {
                        n += 1;
                    }

                    // down
                    x = col;
                    if (y >= 0 and y < nrows and x >= 0 and x < ncols and grid.items[y].items[x] == '@') {
                        n += 1;
                    }

                    // down right
                    x = col + 1;
                    if (y >= 0 and y < nrows and x >= 0 and x < ncols and grid.items[y].items[x] == '@') {
                        n += 1;
                    }

                    if (n < 4) {
                        // std.debug.print("{d}, {d} = {d}\n", .{ row, col, n });
                        grid.items[row].items[col] = 'X';
                        rolls += 1;
                        // break;
                    }
                }
            }
        }
        rolls_removed += rolls;
    }

    // cleanup
    for (0..grid.items.len) |i| {
        grid.items[i].deinit(allocator);
    }
    grid.deinit(allocator);

    std.debug.print("p2: {d}\n", .{rolls_removed});
}
