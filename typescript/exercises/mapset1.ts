/*
Return an array of userIds who have performed more than one unique action, sorted alphabetically.

Requirements:

Actions are case-sensitive

Ignore duplicate actions per user

Return empty array if none qualify
*/
Example:
input = [
  { userId: "u1", action: "login" },
  { userId: "u1", action: "view" },
  { userId: "u1", action: "login" },
  { userId: "u2", action: "login" },
  { userId: "u3", action: "logout" },
  { userId: "u3", action: "logout" }
]

output = ["u1"]