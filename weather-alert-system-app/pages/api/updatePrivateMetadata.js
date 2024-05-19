import { getAuth, clerkClient } from "@clerk/nextjs/server";

export default async function handler(req, res) {
  // Allow only POST and GET methods
  if (req.method !== "POST" && req.method !== "GET") {
    return res.status(405).json({ error: "Method Not Allowed" });
  }

  if (req.method === "GET") {
    // Extract userId from query parameters
    const { userId } = req.query;

    if (!userId) {
      return res.status(400).json({ error: "User ID is required" });
    }

    try {
      // Fetch the user data from Clerk
      const user = await clerkClient.users.getUser(userId);
      // Extract the private metadata
      const privateMetadata = user.privateMetadata;
      // Return the private metadata in the response
      res.status(200).json({ privateMetadata });
    } catch (error) {
      console.error("Error fetching user data:", error);
      res.status(500).json({ error: "Failed to fetch user data" });
    }
  }

  if (req.method === "POST") {
    const { userId } = req.body;

    if (!userId) {
      return res.status(400).json({ error: "User ID is required" });
    }

    try {
      // Update the user's private metadata
      await clerkClient.users.updateUserMetadata(userId, {
        privateMetadata: { subscription: true },
      });
      res
        .status(200)
        .json({ message: "User subscription updated successfully" });
    } catch (error) {
      console.error("Error updating user metadata:", error);
      res
        .status(500)
        .json({ error: "Failed to update user subscription metadata" });
    }
  }
}
