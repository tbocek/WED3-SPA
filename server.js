const Vue = require('vue')
const server = require('express')()
const renderer = require('vue-server-renderer').createRenderer()
const fetch = require('node-fetch');


server.get('*', (req, res) => {
  const app = new Vue({
        data: {
            name: 'Hello Vue!',
            message: '',
            counter: 0,
        },
        created: function () {
            fetch('http://localhost:8080/api')
                .then(response => response.json())
                .then(data => {
                    let string = 'Markets:<br>';
                    for (let s of data) {
                        string += s.name + ', price=' + s.price_usd + '<br>';
                    }
                    this.name = string;
                });
        },
        methods: {
            inc: function () {
                this.counter++
            }
        },
        watch: {
            counter: function (oldValue, newValue) {
                this.message = 'changed from '+oldValue+' to '+newValue
            }
        },
        template: `<div>
    <h2 v-html="name" class="hello-title">Hello {{name}}!</h2>
    <input v-model="message" placeholder="edit me">
    <button @click="inc">Add 1</button>
    <p>The button above has been clicked {{ counter }} times.</p></div>
`
    })
    renderer.renderToString(app, (err, html) => {
		if (err) {
			res.status(500).end('Internal Server Error ' + err)
			return
		}
		res.end(`
			<!DOCTYPE html>
			<html lang="en">
				<head><title>Hello</title></head>
				<body>${html}</body>
			</html>
			`)
		})
})

server.listen(8081)

