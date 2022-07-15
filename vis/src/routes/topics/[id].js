export async function get({ params }) {
    let id = params.id
    return {
        status: 200,
        headers: {},
        body: { id },
    }
}
