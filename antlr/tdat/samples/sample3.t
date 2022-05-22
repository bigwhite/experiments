// samples

r0001: Each { || $speed > 30 } => ("speed", "temperature", "salinity");
r0002: None { |,30| $temperature > 5 } => ("speed", "temperature", "salinity");
r0003: None { |3,| $temperature > 10 } => ("speed", "temperature", "salinity");
r0004: Any { |11,30| ($speed < 5) and ($temperature < 2) and (roundUp($salinity) < 600) } => ("speed", "temperature", "salinity");
r0005: Each { |,| (($speed < 5) and (($temperature + 1) < 10)) or ((roundDown($speed) <= 10) and (roundUp($salinity) >= 500))} => ();
