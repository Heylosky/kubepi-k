import {post, get, del, put} from "@/plugins/request"

const baseUrl = "/api/v1/dmps"

export function searchDmps(page, size, conditions) {
    let url = `${baseUrl}/search?pageNum=${page}&pageSize=${size}&showExtra=true`
    return post(url, {conditions: conditions})
}