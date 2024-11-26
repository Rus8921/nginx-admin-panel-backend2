import React, { useEffect, useState } from "react";
import { PermissionsCard } from "../../ui/components/cards/PermissionsCard";
import { IPermission } from "../../types/permissions";
import nginxPanelApiService from "../../api/NginxPanelApiService";
import { Loader } from "react-feather";

export const PermissionsPage = () => {
  const list = [];
  for (let i = 0; i < 5; i++) {
    const innerList = [];
    for (let a = 0; a < 4; a++) {
      innerList.push(a);
    }
    list.push(innerList);
  }

  const [permissions, setPermissions] = useState<IPermission[]>([]);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    let data = nginxPanelApiService.getAllPermissions();
    data.then((resp) => {
      setPermissions(resp);
      setIsLoading(false);
    });
  }, []);

  return isLoading ? (
    <main className="relative w-full flex flex-col items-center justify-center">
      <Loader className="absolute animate-spin text-main-clr" />
    </main>
  ) : (
    <main className="w-full overflow-auto">
      <div className="p-5 flex flex-col gap-1 min-w-[800px]">
        {permissions.map((item) => {
          return <PermissionsCard data={item} />;
        })}
      </div>
    </main>
  );
};
