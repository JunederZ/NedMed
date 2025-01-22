<script lang="ts">
    import PhotoCard from "$lib/components/PhotoCard.svelte";
    import type { PageData } from "./$types.js";

    let { data }: { data: PageData } = $props()

</script>

<div class="flex flex-col items-center h-full w-screen  text-center text-white">
    <div class="mt-8">
        <h1 class="text-4xl font-bold text-center">Media Archive</h1>
        <p class="text-center">Keep your memories!</p>
    </div>
    <div class="flex flex-wrap justify-center gap-6 py-8">
        {#each data.photoData as photo, index}
            <PhotoCard {...photo} img={data.imgData[index]}/>
        {/each}
    </div>
</div>