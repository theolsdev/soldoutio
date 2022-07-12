import request from "request-promise";
import Functionnality from "../tools/functionnality";

export default class extends Functionnality {
    constructor(product, errorDelay) {
        super(product, errorDelay)
        this.product = product;
        this.errorDelay = errorDelay;
        this.flow();
    }

    /** Initialize Function */
    async initialize() {
        this.jar = request.jar();
        this.request = request.defaults({
            followAllRedirects: true,
            resolveWithFullResponse: true,
            proxy: undefined,
            withCredentials: true,
            strictSSL: false,
        });
    }

    /** Main function Flow */
    async flow() {
        console.log("Adding to cart...");
        let addToCartState = await this.addToCart();
        if (addToCartState == false) {
            // Boucle tant que c'est pas true
        }
        // Si c'est true, on continue...
    }

    /** Adding to Cart */
    async addToCart() {
        try {
            const res = await this.request.post({
                url: "https://www.e.leclerc/api/rest/oms-order-api/cart/compute-local-cart-from-offers",
                headers: {
                    "User-Agent": await this.getUserAgent(),
                    "authority": "www.e.leclerc",
                    "method": "POST",
                    "path": "/api/rest/oms-order-api/cart/compute-local-cart-from-offers",
                    "scheme": "https",
                    "accept": "application/json, text/plain, */*",
                    "accept-language": "fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7",
                    "content-type": "application/json",
                    "origin": "https://www.e.leclerc",
                    "referer": this.product,
                    "sec-ch-ua": `".Not/A)Brand";v="99", "Google Chrome";v="103", "Chromium";v="103"`,
                    "sec-ch-ua-mobile": "?0",
                    "sec-ch-ua-platform": "macOS",
                    "sec-fetch-dest": "empty",
                    "sec-fetch-mode": "cors",
                    "sec-fetch-site": "same-origin"
                },
                // Note: Encore une fois, le payload n'est pas dynamique c'est brute...
                body: JSON.stringify([{
                    "offerId": 2294206,
                    "productSku": 7640305958908,
                    "quantity": 1,
                    "slug": "bottines-en-cuir-a-lacet-outdor",
                    "productLogisticClassCode": "TARIF2",
                    "stock": 8,
                    "isAdProduct": true,
                    "offerPrice": {
                        "price": 5990
                    }
                }])
            });
            // Success return true
            if (res.statusCode == 201) {
                return true;
            } else {
                return false;
            }
        } catch (error) {
            // Catch-up l'erreur
            return false;
        }
    }
    
}