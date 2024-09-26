import { Inter } from "next/font/google";
import "./globals.css";
import { ClerkProvider } from "@clerk/nextjs";
import { Analytics } from "@vercel/analytics/react"

const inter = Inter({ subsets: ["latin"] });

export const metadata = {
  title: "Ryan Marando",
  description: "Get the weather from the future",
};

export default function RootLayout({ children }) {
  return (
    <ClerkProvider>
      <html lang="en">
        <head>
          <link
            rel="icon"
            href="/icon?<generated>"
            type="image/<generated>"
            sizes="<generated>"
          />
          <link
            rel="apple-touch-icon"
            href="/icon?<generated>"
            type="image/<generated>"
            sizes="<generated>"
          />
        </head>
          <body className={inter.className}>{children}
          <Analytics/>
          </body>
      </html>
    </ClerkProvider>
  );
}
