import PhotoCard from "$lib/components/PhotoCard.svelte";
import type {PageServerLoad} from "../../.svelte-kit/types/src/routes/$types.js";
import DatabaseFetcher from "$lib/server/database/fetch.js";

let photos: PhotoCard[] = [];

export const load: PageServerLoad = async () => {

    const fetch = new DatabaseFetcher()

    photos = await fetch.getAllPhotos();
    return {
        photos
    };
    
};