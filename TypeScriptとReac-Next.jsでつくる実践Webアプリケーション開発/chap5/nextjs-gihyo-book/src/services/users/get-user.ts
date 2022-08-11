import { ApiContext, User } from '../../types/data'
import { fetcher } from '../../utils'

export type GetUserParams = {
    id: number
}

/**
 * User API
 * @param context API context
 * @param params paramter
 * @returns user
 */
const getUser = async (
    context: ApiContext,
    { id }: GetUserParams,
): Promise<User> => {
    return await fetcher(
        `${context.apiRootUrl.replace(/\/$/g, '')}/users/${id}}`,
        {
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/json',
            },
        },
    )
}

export default getUser;