/*
Return an array of userIds who have performed more than one unique action, sorted alphabetically.

Requirements:

- Actions are case-sensitive
- Ignore duplicate actions per user
- Return empty array if none qualify

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
*/

export function usersWithMultipleUniqueActions(
  logs: { userId: string; action: string }[],
): string[] {
  let actions: Map<string, Set<string>> = new Map<string, Set<string>>();

  // capture unique actions
  for (let {userId, action} of logs) {
    
    if (!actions.has(userId)){
      actions.set(userId, new Set<string>())
    }
  
    actions.get(userId)!.add(action) // non-null assertion safe after has/set

    
  }

  //console.log(actions);
  let results: string[] = [];

  // get list of items that have more that one action
  for (const [key, value] of actions) {
    if (value.size > 1) {
      results.push(key);
    }
  }

  results.sort();

  return results;
}

const input = [
  { userId: "sam", action: "login" },
  { userId: "sam", action: "view" },
  { userId: "sam", action: "login" },
  { userId: "u2", action: "login" },
  { userId: "todd", action: "logout" },
  { userId: "todd", action: "logout" },
  { userId: "billy", action: "login" },
  { userId: "billy", action: "logout" },
];

//output = ["u1"]
console.log("--- testing ----");
console.log(usersWithMultipleUniqueActions(input));
