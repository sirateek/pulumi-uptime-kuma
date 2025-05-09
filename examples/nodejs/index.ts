import * as pulumi from "@pulumi/pulumi";
import * as uptime-kuma from "@pulumi/uptime-kuma";

const myRandomResource = new uptime-kuma.Random("myRandomResource", {length: 24});
export const output = {
    value: myRandomResource.result,
};
