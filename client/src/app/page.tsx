import { getAccessToken } from "@auth0/nextjs-auth0";
import Login from "../components/auth/login";
import Logout from "../components/auth/logout";
import { useUser } from "@auth0/nextjs-auth0/client";

export default async function Home() {
  return (
    <>
      <p>Hello</p>
    </>
  );
}
