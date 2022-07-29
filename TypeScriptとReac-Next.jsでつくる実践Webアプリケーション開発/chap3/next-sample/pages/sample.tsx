import Link from "next/link";
import { useRouter } from "next/router";

export default function Sample() {
    const router = useRouter();
    return (
        <>
        <button onClick={() => router.back()}>Back</button>
        <button onClick={() => router.reload()}>Reload</button>
        </>
    )
}

