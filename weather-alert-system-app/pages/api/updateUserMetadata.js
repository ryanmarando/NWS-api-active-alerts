import { getAuth, clerkClient } from "@clerk/nextjs/server";

export default async function handler(req, res) {
  if (req.method !== "POST") {
    return res.status(405).json({ error: "Method Not Allowed" });
  }

  const { userId, metadata } = req.body;

  try {
    console.log(userId, metadata);
    await clerkClient.users.updateUserMetadata(userId, {
      publicMetadata: metadata,
    });
    res.status(200).json({ message: "User metadata updated successfully" });
  } catch (error) {
    console.error("Error updating user metadata:", error);
    res.status(500).json({ error: "Failed to update user metadata" });
  }
}
