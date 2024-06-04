import { useContext, useState } from "react"
import { APIServer } from "../App"
import { TokenContext } from "../context"


export const useAPI = (callback) => {
    const [isLoading, setIsLoading] = useState(false)
    const [error, setError] = useState("")
    const { setToken } = useContext(TokenContext)

    const featching = async () => {
        try {
            setIsLoading(true)
            await callback()
        } catch (e) {
            setError(e)
            if (e.response !== undefined && [401, 403].some(code => code === e.response.status)) {
                localStorage.setItem(APIServer.tokenName, "")
                setToken("")
            }
        } finally {
            setIsLoading(false)
        }
    }

    return [featching, isLoading, error]

}