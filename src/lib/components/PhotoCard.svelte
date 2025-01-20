<script lang="ts">

    import NoImg from "../assets/no-img.png"

    let props = $props();
    console.log(props);

    const downloadImage = () => {
        if (props.url) {
            const link = document.createElement('a');
            link.href = props.url;
            link.download = props.title || 'downloaded_image';
            link.style.display = 'none'; 
            document.body.appendChild(link);
            link.click();
            document.body.removeChild(link);
        }
    };

</script>


<div class="flex flex-col h-full bg-gray-600/60 p-4 rounded-xl">
    <div>
        {#if props.filename == null || props.filename === ""}
            <h1>No Title</h1>
        {:else}
            <h1>{props.filename}</h1>
        {/if}
    </div>
    <div>
        {#if props.url == null || props.url === ""}
            <img src={NoImg} alt="" class="w-auto h-40 aspect-square object-cover" />
        {:else}
            <img src={props.url} alt="" class="w-auto h-40 aspect-square object-cover" />
        {/if}
    </div>
    <div>
        {#if props.description == null || props.description === ""}
            <p>No Description.</p>
        {:else}
            <p>{props.description}</p>
        {/if}
        <div>
            <button onclick={downloadImage} disabled={!props.url}>Download</button>
        </div>
    </div>
</div>