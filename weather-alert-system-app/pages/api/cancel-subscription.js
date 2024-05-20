// pages/api/cancel-subscription.js
import Stripe from "stripe";

const secretTESTKEY =
  "sk_test_51PHaB5BhEzKZHebbt9ppEiox7yIs2G3gtw5gIELUvjNkpy2hgCl5C71oiqvKO2EcePRv0iO8tITAwxU979qmDULI006swoFnef";

const stripe = new Stripe(secretTESTKEY, {
  apiVersion: "2024-04-10",
});

export default async function handler(req, res) {
  if (req.method !== "POST") {
    return res.status(405).send({ message: "Only POST requests allowed" });
  }

  const { subscriptionId } = req.body;

  if (!subscriptionId) {
    return res.status(400).send({ message: "Subscription ID is required" });
  }

  try {
    // Cancel the subscription
    const deletedSubscription = await stripe.subscriptions.cancel(
      subscriptionId
    );

    // Respond with the result of the cancellation
    res.status(200).json({ subscription: deletedSubscription });
  } catch (error) {
    console.error("Error canceling subscription:", error);
    res.status(500).send({ error: error.message });
  }
}
