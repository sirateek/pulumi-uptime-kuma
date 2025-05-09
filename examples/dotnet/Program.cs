using System.Collections.Generic;
using System.Linq;
using Pulumi;
using uptime-kuma = Pulumi.uptime-kuma;

return await Deployment.RunAsync(() => 
{
    var myRandomResource = new uptime-kuma.Random("myRandomResource", new()
    {
        Length = 24,
    });

    return new Dictionary<string, object?>
    {
        ["output"] = 
        {
            { "value", myRandomResource.Result },
        },
    };
});

