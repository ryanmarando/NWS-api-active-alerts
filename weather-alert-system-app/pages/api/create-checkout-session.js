import Stripe from "stripe";

const secretTESTKEY =
  "sk_test_51PHaB5BhEzKZHebbt9ppEiox7yIs2G3gtw5gIELUvjNkpy2hgCl5C71oiqvKO2EcePRv0iO8tITAwxU979qmDULI006swoFnef";

const stripe = new Stripe(secretTESTKEY, {
  apiVersion: "2024-04-10",
});

export default async function handler(req, res) {
  if (req.method === "POST") {
    try {
      const session = await stripe.checkout.sessions.create({
        payment_method_types: ["card"],
        line_items: [
          {
            price: "price_1PHbNUBhEzKZHebb7MYPxyvG",
            quantity: 1,
          },
        ],
        mode: "subscription",
        success_url: `${req.headers.origin}/success`,
        cancel_url: `${req.headers.origin}/`,
      });
      res.status(200).json({ id: session.id });
    } catch (err) {
      console.error("Error creating Stripe session:", err);
      res.status(500).json({ error: err.message });
    }
  } else {
    res.setHeader("Allow", "POST");
    res.status(405).end("Method Not Allowed");
  }
}
