Ejercicios para afianzar el patrón

1. Agrega una Radio que implemente Device y úsala con los mismos OnCommand y OffCommand. Vas a ver la reutilización en acción.
2. Crea un ToggleCommand que encienda o apague según isRunning. Tendrás que exponer el estado en Device.
3. Implementa undo: agrega undo() a la interfaz Command (OnCommand.undo() llama a off()) y un RemoteControl que guarde el historial de comandos ejecutados en un slice y tenga un método Undo(). Este ejercicio muestra la verdadera ventaja del patrón.
4. Macro-comando: un MacroCommand que contenga un []Command y los ejecute todos, por ejemplo "modo cine": apagar las luces y encender la TV.
