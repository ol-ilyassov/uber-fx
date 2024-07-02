
# Notes

The work starts with **fx.New()**.

**fx.Run()** | **fx.Start()** | **fx.Stop()**

**fx.Provide()** - provides values to the container and make them available to the application.

**fx.Invoke()** - used for root-level invocations (starting a server, running a main loop, etc.), running functions that have side effects (starting a background worker, configuring a global logger, etc.).

**fx.Lifecycle** contains 2 phases:

[1] **Initialization**:
- register all constructors passed to **fx.Provide**;
- reigster all decorators passed to **fx.Decorate**;
- run all functions passed to **fx.Invoke**, calling constructors and decorators as needed.

[2] **Execution**:
- run all **startup hooks** appended to the application by providers, decorators, and invoked functions;
- wait for a signal to stop running;
- run all **shutdown hooks** appended to the application.

Further: 
- **fx.Module**;
- **Parameter Object**;
- **Result Object**;
- **Annotations** / **Value Groups**.
