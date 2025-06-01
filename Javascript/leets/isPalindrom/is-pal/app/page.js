"use client";
import Image from "next/image";
import styles from "./page.module.css";
import { useState } from 'react';



export default function Home() {
   const [input, setInput] = useState('');
  const [result, setResult] = useState(null);

  function isPalindrome(val) {
    let stack = [];
    let value = val;

    while (value > 0) {
      let digit = value % 10;
      stack.push(digit);
      value = Math.floor(value / 10);
    }

    let rev = 0;
    for (let i = 0; i < stack.length; i++) {
      rev = rev * 10 + stack[i];
    }

    return rev === val;
  }

  function handleCheck() {
    const num = parseInt(input, 10);
    if (!isNaN(num)) {
      setResult(isPalindrome(num));
    } else {
      setResult(null);
    }
  }

  return (
    <div style={{ padding: '2rem' }}>
      <h1>Palindrome Checker</h1>
      <input
        type="text"
        value={input}
        onChange={(e) => setInput(e.target.value)}
        placeholder="Enter a number"
      />
      <button onClick={handleCheck}>Check</button>
      {result !== null && (
        <p>{result ? '✅ It is a palindrome!' : '❌ Not a palindrome.'}</p>
      )}
    </div>
  );
}
