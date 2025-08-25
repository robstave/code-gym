/**
 * Draw a simple text triangle to the console.
 *
 * Usage:
 *  - triangle(5)                        // left-aligned, '*'
 *  - triangle(5, { align: 'right' })    // right-aligned
 *  - triangle(5, { align: 'center', char: '#' })
 *
 * @param {number} size  Number of rows (>= 0)
 * @param {{align?: 'left'|'right'|'center', char?: string}} [options]
 */
function triangle(size, options = {}) {
  const { align = "left", char = "*" } = options;

  const rows = Math.max(0, Number.isFinite(size) ? Math.trunc(size) : 0);
  const c = typeof char === "string" && char.length > 0 ? char[0] : "*";

  for (let i = 1; i <= rows; i++) {
    const block = c.repeat(i);
    let line;

    if (align === "right") {
      line = " ".repeat(rows - i) + block;
    } else if (align === "center") {
      const left = Math.floor((rows - i) / 2);
      const right = rows - i - left;
      line = " ".repeat(left) + block + " ".repeat(right);
    } else {
      // default 'left'
      line = block;
    }

    console.log(line);
  }
}

// Example run (kept to match original behavior)
triangle(5);

triangle(5, { align: "right" }); // right-aligned
triangle(5, { align: "center", char: "#" });
