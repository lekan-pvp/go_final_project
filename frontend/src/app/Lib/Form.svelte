<script>
    import FormCombo from "./FormCombo.svelte";
    import FormInput from "./FormInput.svelte";
    import FormCheck from "./FormCheck.svelte";
    import FormText from "./FormText.svelte";

    export let ctrls = {};
    export let values = {};
    export let linked = {};
    export let callback;

    function changeValue(field) {
        callback("change", field);
    }
</script>

{#each ctrls as ctrl}
    {#if ctrl.field}
        {#if ctrl.type == "input" || ctrl.type == "date" || ctrl.type == "money"}
            <FormInput
                type={ctrl.type}
                title={ctrl.title}
                width={ctrl.width}
                size={ctrl.size}
                req={ctrl.req}
                change={() => {
                    changeValue(ctrl.field);
                }}
                bind:value={values[ctrl.field]}
            />
        {:else if ctrl.type == "text"}
            <FormText
                title={ctrl.title}
                size={ctrl.size}
                change={() => {
                    changeValue(ctrl.field);
                }}
                bind:value={values[ctrl.field]}
            />
        {:else if ctrl.type == "check"}
            <FormCheck
                title={ctrl.title}
                width={ctrl.width}
                change={() => {
                    changeValue(ctrl.field);
                }}
                bind:value={values[ctrl.field]}
            />
        {:else if ctrl.type == "link"}
            <FormInput
                type={ctrl.type}
                title={ctrl.title}
                width={ctrl.width}
                size={ctrl.size}
                req={ctrl.req}
                link={ctrl.link}
                linkedit={ctrl.linkedit}
                change={() => {
                    changeValue(ctrl.field);
                    if (linked[ctrl.field]) {
                        delete linked[ctrl.field];
                    }
                }}
                callback={(action, par) => {
                    changeValue(ctrl.field);
                    values[ctrl.field] = par.name;
                    linked[ctrl.field] = par.id;
                }}
                bind:value={values[ctrl.field]}
            />
        {:else if ctrl.type == "combo"}
            <FormCombo
                req={ctrl.req}
                title={ctrl.title}
                width={ctrl.width}
                change={() => {
                    changeValue(ctrl.field);
                }}
                list={ctrl.list}
                bind:value={values[ctrl.field]}
            />
        {/if}
    {/if}
{/each}
