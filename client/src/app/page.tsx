import { getAccessToken } from "@auth0/nextjs-auth0";
import Login from "./components/auth/login";
import Logout from "./components/auth/logout";

export default async function Home() {
    const { accessToken } = await getAccessToken();
    return (
        <>
            <Login />
            <Logout />
            <p>Hello</p>
            {accessToken}
        </>
    );
}
