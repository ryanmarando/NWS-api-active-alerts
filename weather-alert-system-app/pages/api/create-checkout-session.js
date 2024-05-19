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
        success_url: `${req.headers.origin}/success/3Pb0xmpcABuhT8DpZZU7GMLCExF1is9nsfhb3zrWmGTwIGuAGB7X4ODZwMtRM8iSJ5Y1uPt1OUXAuJ0MqJ4LfOC0P4PW22zfy5HhX3eV5cmaQBTVXofyLg4lXCAaE7BIrQsk4tpsFyPlmWRAPWEK73txs6P5g3B4e654TIn0stKv8gzgzGRdfYcko3KlnUrnA0VtLi1c2vmAZhBTEoPBQUcmn3uWf1tohjnEX9AQyEbEhY3EAMx9L9XyIN3GlErwEHl5AQVZbID9yVWyZCzmohFru9kU4yZsc3iEMj7uaRobUyHpSJJrC6BKEWYOVYR7TsQXEtzsCOaKlwrHDbXwgPfxX3k4rQIXbkwjb0jAJqiIqg3vrX1jMcTALt2uuSk1WofpiJo9Mo9HjHmRHHkPxWLZRLrMVXZL47W5kgOQEifizWGkWbQDgBj36bllSuhHbM6jZhSSXRZw9qcbNslHn1uoOm8g2kf8s3HQzq2ocI2Yw1EncUT5qz1SIxEBoC0qNBLCXEKcnDwrbbUxxRafsM9qGwKs90Uid3CUotY23AMRwxXgmIKejUDuZJ9aRCu911PUDd63cKSXvhIE3VSnir0hLDNoRdCFOw7qtyTdGitd0pqQvz5CHTTAljmaexQUBEQqEQZCZt3AvQkda4yCImNRZn9l3vInjgnIGnLDj5edq5yfY1aaK6KhP06L9anqNwIkM0Jvkw6OM9gUU3pYC97e9urdYgtJrIYStTJEB4DGH7t5132kqfFnn67aE6nhWsI0gT2tZZN53DTZYSa3wo4KT`,
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
