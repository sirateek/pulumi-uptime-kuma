import pulumi
import pulumi_uptime-kuma as uptime-kuma

my_random_resource = uptime-kuma.Random("myRandomResource", length=24)
pulumi.export("output", {
    "value": my_random_resource.result,
})
