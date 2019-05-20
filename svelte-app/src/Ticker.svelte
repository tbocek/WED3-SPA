<style>
    /* nice background color */
    body {
        background-color:#f0f0f0;
    }
</style>

<script>
    import { onMount } from 'svelte';
	export let symbol;

	let coinmarketcapData = {};
	let price;

	onMount(async () => {
	    console.log("22AA"+coinmarketcapData);
        const res = await fetch('https://api.coinmarketcap.com/v1/ticker/?limit=10');
        let body = await res.json();

        coinmarketcapData = body.reduce(function(map, obj) {
            map[obj.symbol] = obj;
                return map;
            }, {});
        price = calcPrice();
    });

	function calcPrice() {
	    if (typeof coinmarketcapData[symbol] !== 'undefined') {
	        console.log(coinmarketcapData);
            return coinmarketcapData[symbol].name + " " + coinmarketcapData[symbol].price_usd;
        }
    }
</script>

<p>{price} USD</p>