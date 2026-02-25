package main

import (
	"fmt"
	"log"

	"github.com/codiume/termwind"
)

func main() {
	output, err := termwind.Render(`
        <div class="ml-2">
            <span class="px-1 mt-1 bg-green-500 text-white font-bold">⚙️ main</span>
            <span class="px-1 mt-1 bg-blue-500 text-white">v1.0.0</span>
            <span class="ml-1">Style CLI applications with tailwind like syntax</span>
            <hr />
            <ul>
                <li class="text-red-400">Fast</li>
                <li class="text-yellow-400">Simple and easy</li>
                <li class="text-green-400">Composable Elements</li>
                <li class="text-cyan-400">Customizable</li>
            </ul>
            <dl>
                <dt>Author</dt>
                <dd>Jane Doe</dd>
                <dt>License</dt>
                <dd>MIT</dd>
            </dl>
            <a href="https://github.com/codiume/termwind">View on GitHub</a>
        </div>
    `)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(output)
}
